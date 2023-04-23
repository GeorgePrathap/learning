package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	// Create a simple HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)

		// Write some sample data to the response body
		w.Write([]byte("Hello, world!"))
	})

	// Create a new CompressHandler with default options
	compressHandler := handlers.CompressHandler(handler)

	// Create a new CompressHandler with custom compression level
	// compressHandler := handlers.CompressHandlerLevel(handler, 6)

	// Create a new CompressHandler with custom minimum content length for compression
	// compressHandler := handlers.CompressHandler(handler, handlers.CompressionMinSize(1024))

	// Create a new CompressHandler with custom set of content types to exclude from compression
	// compressHandler := handlers.CompressHandler(handler, handlers.ExcludedContentTypes([]string{"image/jpeg", "image/png"}))

	// Create a new server using the CompressHandler
	server := &http.Server{
		Addr:    ":8080",
		Handler: compressHandler,
	}

	// Start the server
	server.ListenAndServe()
}

/*

	handlers.CompressHandler(handler):
		creates a new CompressHandler middleware with default options and passes the HTTP handler handler to it. This means that any requests to the server will be handled by the CompressHandler middleware, which will automatically compress the response if the client supports gzip compression.

	handlers.CompressHandlerLevel(handler, level):
	 	creates a new CompressHandler middleware with the specified compression level and passes the HTTP handler handler to it. The level argument must be an integer between 0 and 9, with 0 indicating no compression and 9 indicating maximum compression. The default compression level is 6.

	handlers.CompressHandler(handler, handlers.CompressionMinSize(size)):
		creates a new CompressHandler middleware with the specified minimum content length for compression and passes the HTTP handler handler to it. The size argument specifies the minimum content length for compression, in bytes. Responses with a content length smaller than size will not be compressed. The default minimum content length is 1024 bytes.

		The size argument specifies the minimum content length for compression, in bytes. This means that the CompressHandler middleware will only compress HTTP responses that have a content length larger than or equal to size. For example, if size is set to 1024 bytes, then only responses with a content length of 1024 bytes or larger will be compressed.

		Why do we do this? Compressing small amounts of data can sometimes result in a larger compressed file size due to the overhead of the compression algorithm. So, it's not always efficient to compress small files. That's why we set a minimum content length for compression. If the content length of the response is smaller than the minimum size, the response will not be compressed.

	handlers.CompressHandler(handler, handlers.ExcludedContentTypes(types)):
		 creates a new CompressHandler middleware with the specified set of content types to exclude from compression and passes the HTTP handler handler to it. The types argument is a slice of MIME types that should not be compressed. The default excluded content types are image/*, audio/*, and video/*.

*/
