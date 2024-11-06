package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

// Data structure to represent blockchain data
type BlockchainData struct {
    Seq      uint64 `json:"seq"`
    HashCode string `json:"hashcode"`
    Status   string `json:"status"` // e.g., "Shipped" or "Sold"
}

// A simple in-memory store to simulate blockchain data
var blockchain = make(map[uint64]BlockchainData)
var mu sync.Mutex // Mutex for thread-safe access to blockchain map

// Handler to store data on blockchain (simulates writing to blockchain)
func storeDataHandler(w http.ResponseWriter, r *http.Request) {
    var data []BlockchainData
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    mu.Lock()
    for _, entry := range data {
        // Store data in memory
        blockchain[entry.Seq] = entry
        fmt.Printf("Storing data on blockchain: Seq=%d, HashCode=%s\n", entry.Seq, entry.HashCode)
    }
    mu.Unlock()

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"success"}`))
}

// Handler to retrieve a specific record or all records from blockchain
func getRecordsHandler(w http.ResponseWriter, r *http.Request) {
    seq := r.URL.Query().Get("seq")

    mu.Lock()
    defer mu.Unlock()

    if seq != "" {
        // Parse `seq` parameter if provided
        var seqUint uint64
        _, err := fmt.Sscan(seq, &seqUint)
        if err != nil {
            http.Error(w, "Invalid seq parameter", http.StatusBadRequest)
            return
        }

        // Retrieve specific record if it exists
        record, exists := blockchain[seqUint]
        if !exists {
            http.Error(w, "Record not found", http.StatusNotFound)
            return
        }

        // Return JSON response with specific record
        json.NewEncoder(w).Encode(record)
    } else {
        // Return all records if no specific `seq` is provided
        records := make([]BlockchainData, 0, len(blockchain))
        for _, record := range blockchain {
            records = append(records, record)
        }
        json.NewEncoder(w).Encode(records)
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/storeData", storeDataHandler)
    mux.HandleFunc("/getRecords", getRecordsHandler)

    // CORS middleware setup
    handler := corsMiddleware(mux)

    log.Println("Starting Go server on port 8081...")
    log.Fatal(http.ListenAndServe(":8081", handler))
}

// Middleware to enable CORS
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}
