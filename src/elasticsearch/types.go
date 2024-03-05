package elasticsearch

const (
	esAddr = "http://localhost:9200"
)

const createIndexBody = `{
    "settings": {
        "number_of_shards": 3,
        "number_of_replicas": 2
    },
    "mappings": {
        "properties": {
            "name": {
                "type": "keyword"
            },
            "email": {
                "type": "text"
            },
            "phone": {
                "type": "keyword"
            },
            "birthday": {
                "type": "date"
            }
        }
    }
}`

type User struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Birthday int64  `json:"birthday"`
}
