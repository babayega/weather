package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

var (
	mu        sync.Mutex
	lastCalls = make(map[string]time.Time)
)

func WeatherRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.GetHeader("X-Message")

		// Lock mutex to get and update lastCalls map
		mu.Lock()
		lastCallTime, exists := lastCalls[address]
		lastCalls[address] = time.Now()
		mu.Unlock()

		if exists && time.Since(lastCallTime) < 12*time.Second {
			// Client has made a recent call, check if it's within the first 2 seconds of the interval
			if time.Since(lastCallTime) <= 2*time.Second {
				// Allow the request to proceed
				c.Next()
				return
			}

			// Request made outside the allowed timeframe, return an error response
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "You can only report the weather once every 12 seconds, within the first 2 seconds of the interval",
			})
			c.Abort()
			return
		}

		// First request from the client within the interval, allow it to proceed
		c.Next()
	}
}

func VerifySignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("Incoming request:", c.Request.URL.Path)

		// Read the signed message and signature from the request headers
		message := c.GetHeader("X-Message")
		signature := c.GetHeader("X-Signature")
		publicKeyHex := message

		// Verify the public key from the signed message
		isValid := verifyPublicKey(message, signature, publicKeyHex)

		if !isValid {
			// Return a 401 Unauthorized response if the public key check fails
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Pass the request to the next middleware/handler
		c.Next()

		println("Response status code:", c.Writer.Status())
	}
}

func verifyPublicKey(message, signature, publicKeyHex string) bool {
	// Get the hash of the message
	// "Ethereum Signed Message" is appended by web3/ethers libraries when generating
	// a signed message. Hence need to add it here to verify the message.
	msgBytes := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(msgBytes))

	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		println("signatureBytes error: ", err)
		return false
	}

	publicKeyBytes, err := hexutil.Decode(publicKeyHex)
	if err != nil {
		println("publicKeyBytes error: ", err)
		return false
	}

	return crypto.VerifySignature(publicKeyBytes, messageHash.Bytes(), signatureBytes[:len(signatureBytes)-1])
}
