package main

import (
	"fmt"

	flourite "go-flourite"
)

func main() {
	snippet := `FROM node:16.6-buster

WORKDIR /app

COPY . .

RUN npm install --production

EXPOSE 8080

CMD ["npm", "start"]`

	detector := flourite.NewDetector()
	res := detector.Detect(snippet)
	fmt.Println(res)
	fmt.Println(res.Best())
}
