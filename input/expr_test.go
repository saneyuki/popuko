package input

import (
	"testing"
)

func TestParseCommand1(t *testing.T) {
	ok, cmd := ParseCommand("@bot r+")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if actual := v.priority; actual != 0 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand2(t *testing.T) {
	ok, cmd := ParseCommand("@reviewer r?")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AssignReviewerCommand)
	if !ok {
		t.Fatal("should be AssignReviewerCommand")
	}

	if v.Reviewer != "reviewer" {
		t.Fatal("should be the expected reviewer")
	}
}

func TestParseCommand3(t *testing.T) {
	ok, cmd := ParseCommand("@bot r=popuko,pipimi")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByOthersCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByOthersCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if name := v.Reviewer[0]; name != "popuko" {
		t.Fatalf("should be the expected reviewer 1: %v\n", name)
	}

	if name := v.Reviewer[1]; name != "pipimi" {
		t.Fatalf("should be the expected reviewer 2: %v\n", name)
	}

	if actual := v.priority; actual != 0 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand5(t *testing.T) {
	ok, cmd := ParseCommand("")
	if ok {
		t.Fatal("should be no result")
	}

	if cmd != nil {
		t.Fatal("command should be nil")
	}
}

func TestParseCommand6(t *testing.T) {
	ok, _ := ParseCommand(`@bot
    r+`)
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand7(t *testing.T) {
	ok, _ := ParseCommand("@bot")
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand8(t *testing.T) {
	ok, _ := ParseCommand("bot r+")
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand9(t *testing.T) {
	ok, _ := ParseCommand("Hello, I'm john.")
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand10(t *testing.T) {
	ok, _ := ParseCommand("    @bot r+")
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand11(t *testing.T) {
	ok, _ := ParseCommand(`
    @bot r+`)
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestParseCommand12(t *testing.T) {
	ok, cmd := ParseCommand("r? @reviewer")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AssignReviewerCommand)
	if !ok {
		t.Fatal("should be AssignReviewerCommand")
	}

	if v.Reviewer != "reviewer" {
		t.Fatal("should be the expected reviewer")
	}
}

func TestParseCommand13(t *testing.T) {
	ok, cmd := ParseCommand("@bot r-")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*RejectChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be RejectChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}
}

func TestParseCommand14(t *testing.T) {
	ok, cmd := ParseCommand("@bot r+ p=-1")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if actual := v.priority; actual != -1 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand15(t *testing.T) {
	ok, cmd := ParseCommand("@bot r+ p=0")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if actual := v.priority; actual != 0 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand16(t *testing.T) {
	ok, cmd := ParseCommand("@bot r+ p=1")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if actual := v.priority; actual != 1 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand17(t *testing.T) {
	ok, cmd := ParseCommand("@bot r+ p=+2")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByReviewerCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByReviewerCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if actual := v.priority; actual != 2 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand18(t *testing.T) {
	ok, cmd := ParseCommand("@bot r=popuko,pipimi r=0")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByOthersCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByOthersCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if name := v.Reviewer[0]; name != "popuko" {
		t.Fatalf("should be the expected reviewer 1: %v\n", name)
	}

	if name := v.Reviewer[1]; name != "pipimi" {
		t.Fatalf("should be the expected reviewer 2: %v\n", name)
	}

	if actual := v.priority; actual != 0 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand19(t *testing.T) {
	ok, cmd := ParseCommand("@bot r=popuko,pipimi r=-1")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByOthersCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByOthersCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if name := v.Reviewer[0]; name != "popuko" {
		t.Fatalf("should be the expected reviewer 1: %v\n", name)
	}

	if name := v.Reviewer[1]; name != "pipimi" {
		t.Fatalf("should be the expected reviewer 2: %v\n", name)
	}

	if actual := v.priority; actual != -1 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}

func TestParseCommand20(t *testing.T) {
	ok, cmd := ParseCommand("@bot r=popuko,pipimi r=+2")
	if !ok {
		t.Fatal("should be ok")
	}

	v, ok := cmd.(*AcceptChangeByOthersCommand)
	if !ok {
		t.Fatal("should be AcceptChangeByOthersCommand")
	}

	if name := v.BotName(); name != "bot" {
		t.Fatalf("should be the expected bot name: %v\n", name)
	}

	if name := v.Reviewer[0]; name != "popuko" {
		t.Fatalf("should be the expected reviewer 1: %v\n", name)
	}

	if name := v.Reviewer[1]; name != "pipimi" {
		t.Fatalf("should be the expected reviewer 2: %v\n", name)
	}

	if actual := v.priority; actual != 2 {
		t.Fatalf("should be the expected priority: %v\n", actual)
	}
}
