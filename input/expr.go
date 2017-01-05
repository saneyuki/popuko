package input

import (
	"log"
	"strings"
)

// ParseCommand is doing adhoc command parsing.
// for the future, we should write an actual parser.
func ParseCommand(raw string) (ok bool, cmd interface{}) {
	log.Printf("debug: input: %v\n", raw)
	tmp := strings.Split(raw, "\n")

	// If there are no possibility that the comment body is not formatted
	// `@botname command`, stop to process.
	if len(tmp) < 1 {
		return false, nil
	}

	body := tmp[0]
	log.Printf("debug: body: %v\n", body)

	command := strings.Split(body, " ")
	log.Printf("debug: command: %#v\n", command)
	if len(command) < 2 {
		return false, nil
	}

	trigger := command[0]
	if (strings.Index(trigger, "@") != 0) && (trigger != "r?") {
		return false, nil
	}

	args := strings.Trim(command[1], " ")
	log.Printf("debug: trigger: %v\n", trigger)
	log.Printf("debug: args: %#v\n", args)

	if trigger == "r?" {
		return true, &AssignReviewerCommand{
			Reviewer: strings.TrimPrefix(args, "@"),
		}
	}

	if args == "r?" {
		return true, &AssignReviewerCommand{
			Reviewer: strings.TrimPrefix(trigger, "@"),
		}
	}

	if args == "r-" {
		return true, &RejectChangeByReviewerCommand{
			botName: strings.TrimPrefix(trigger, "@"),
		}
	}

	if args == "r+" {
		return true, &AcceptChangeByReviewerCommand{
			botName: strings.TrimPrefix(trigger, "@"),
		}
	}

	if strings.Index(args, "r=") != 0 {
		return false, nil
	}

	args = strings.TrimPrefix(args, "r=")
	log.Printf("debug: args: %#v\n", args)
	reviwers := strings.Split(args, ",")
	log.Printf("debug: reviwers: %#v\n", reviwers)

	for i, name := range reviwers {
		reviwers[i] = strings.Trim(name, " ")
	}

	return true, &AcceptChangeByOthersCommand{
		botName:  strings.TrimPrefix(trigger, "@"),
		Reviewer: reviwers,
	}
}

type AcceptChangesetCommand interface {
	BotName() string
}

type AcceptChangeByReviewerCommand struct {
	botName string
}

func (s *AcceptChangeByReviewerCommand) BotName() string {
	return s.botName
}

type AcceptChangeByOthersCommand struct {
	botName  string
	Reviewer []string
}

func (s *AcceptChangeByOthersCommand) BotName() string {
	return s.botName
}

type AssignReviewerCommand struct {
	Reviewer string
}

type RejectChangeByReviewerCommand struct {
	botName string
}

func (s *RejectChangeByReviewerCommand) BotName() string {
	return s.botName
}
