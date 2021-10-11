truncate table public.users ;
ALTER SEQUENCE users_seq RESTART WITH 1


func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

--------

func authError(w http.ResponseWriter, err error, clientMsg string) {
    log.Println("Authentication failed: %v", err)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusForbidden)

    err = json.NewEncoder(w).Encode(struct {
        Error string
    }{Error: clientMsg})
    if err != nil {
        log.Println("Failed to write response: %v", err)
    }

    return
}

--------------

created_at ve updated_at -> ISO8601(UTC)