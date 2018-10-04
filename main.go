package main

import (
	"log"

	"github.com/Seanstoppable/gmailcleaner/auth"
	rules "github.com/Seanstoppable/gmailcleaner/rules"
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
)

func main() {
	ctx := context.Background()

	client := auth.GetClient(ctx)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}

	user := "me"

	rules, err := rules.LoadRules()

	if err != nil {
		log.Fatalf("Unable to load config %v", err)
	}

	for _, rule := range rules {
		query := rule.Query.CreateQuery()

		msgs := []*gmail.Message{}
		pageToken := ""
		for {
			req := srv.Users.Messages.List(user).Q(query)
			if pageToken != "" {
				req.PageToken(pageToken)
			}
			r, err := req.Do()
			if err != nil {
				log.Fatalf("Unable to retrieve messages: %v", err)
			}

			log.Printf("Processing %v messages for query %s\n", len(r.Messages), query)
			for _, m := range r.Messages {
				msg, err := srv.Users.Messages.Get("me", m.Id).Do()
				if err != nil {
					log.Fatalf("Unable to retrieve message %v: %v", m.Id, err)
				}
				msgs = append(msgs, msg)
			}
			if r.NextPageToken == "" {
				break
			}
			pageToken = r.NextPageToken
		}
		for _, m := range msgs {
			log.Printf("%s %v", m.Snippet, m.LabelIds)
		}

		if len(rule.Modifications.AddLabels) > 0 || len(rule.Modifications.RemoveLabels) > 0 {

			modificationRequest := &gmail.ModifyMessageRequest{
				AddLabelIds:    rule.Modifications.AddLabels,
				RemoveLabelIds: rule.Modifications.RemoveLabels,
			}

			for _, m := range msgs {
				archiveReq := srv.Users.Messages.Modify("me", m.Id, modificationRequest)
				_, err := archiveReq.Do()
				if err != nil {
					log.Fatalf("Unable to archive messages: %v", err)
				}
			}
		}
	}
}
