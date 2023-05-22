package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type KapetaLogEntry struct {
	NameAndVersion string        `json:"nameAndVersion"`
	Created        time.Time     `json:"created"`
	StateType      string        `json:"stateType"`
	State          KapetaLogLine `json:"state"`
}

type KapetaLogLine struct {
	Message string `json:"message"`
}

const (
	kapetaCallbackEnv = "KAPETA_CALLBACK"
	kapetaTokenEnv    = "KAPETA_CREDENTIALS_TOKEN"
)

func main() {
	kapetaToken := os.Getenv(kapetaTokenEnv)
	if kapetaToken == "" {
		fmt.Printf("%s not defined\n", kapetaTokenEnv)
		os.Exit(1)
	}

	kapetaCallback := os.Getenv(kapetaCallbackEnv)
	if kapetaCallback == "" {
		fmt.Printf("%s not defined\n", kapetaCallbackEnv)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		logLine := scanner.Text()

		kapetaLog := KapetaLogEntry{
			Created:   time.Now(),
			StateType: "plain",
			State:     KapetaLogLine{Message: logLine},
		}

		payload, err := json.Marshal(kapetaLog)
		if err != nil {
			panic(err)
		}
		log.Println(string(payload))
		err = post(payload, kapetaToken, kapetaCallback)
		if err != nil {
			log.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func post(payloadBytes []byte, kapetaToken string, kapetaCallback string) error {

	url := fmt.Sprintf("%s/info/log", kapetaCallback)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", kapetaToken))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("POST %s responded with status: %s", kapetaCallback, resp.Status)
	}
	return nil
}
