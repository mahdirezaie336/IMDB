package handler

func makeResponse(s string) map[string]string {
	switch s {
	case "bad-request":
		return map[string]string{
			"message": "Bad Request",
		}

	case "server-error":
		return map[string]string{
			"message": "Internal Server Error",
		}

	case "ok":
		return map[string]string{
			"message": "OK",
		}

	case "id-not-found":
		return map[string]string{
			"message": "ID Not Found",
		}

	default:
		return map[string]string{}
	}
}
