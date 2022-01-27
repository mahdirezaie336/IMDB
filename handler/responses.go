package handler

func makeResponse(s string) map[string]string {
	switch s {
	case "bad-request":
		return map[string]string{
			"message": "Bad Request!",
		}
	default:
		return map[string]string{}
	}
}
