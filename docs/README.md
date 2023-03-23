If written answers are required, you can add them to this file. Just copy the
relevant questions from the root of the repo, preferably in
[Markdown](https://guides.github.com/features/mastering-markdown/) format :)

#### Task 1

##### Buggy Code 1
1. What is wrong: the main goroutine is asleep because sending "hello world" through channel blocks the main goroutine, preventing the next line from receiving anything through the channel
2. How it was fixed: make the channel buffered, so that it doesn't block (until it reaches > 1 string)

##### Buggy Code 2
1. What is wrong: the main goroutine closes before 11 can be printed
2. How it was fixed: force the main goroutine to wait for input from a channel, which only comes when the printing is done

#### Task 2

| Question | What I expected | What happened | Why I believe this happened |
|-|-|-|-|
| What happens if you do X? |  Program would still work as before | Program ended up in a deadlock | Because of reasons 🤷 |
| What happens if you switch the order of the statements `wgp.Wait()` and `close(ch)` in the end of the `main` function? | an error or funky behavior (like perma-blocking/deadlocking) since we're waiting for input from a closed channel | error | ~~cuz i said so~~because we tried to send to a closed channel |
| What happens if you move the `close(ch)` from the `main` function and instead close the channel in the end of the function `Produce`?  | same as above, since we're closing the channel probably before printing everything | same as above, but at a random time since we don't know which call to Produce will run first | same as above |
| What happens if you remove the statement `close(ch)` completely?  | nothing | nothing | what exactly is supposed to happen? we've reached the end of the program so the channel closes anyway |
| What happens if you increase the number of consumers from 2 to 4?  | nothing special | runtime halved| we have 4 goroutines consuming instead of just 2 |
| Can you be sure that all strings are printed before the program stops?  | no, since we never wait for the consumers  | yes? | since ch is unbuffered, sending something through it makes the Produce fn block until the int is received, so all (except maybe one?) of the strings have to be printed |