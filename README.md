mail-to-gh
==========

Take an email message from stdin and generate a GitHub issue.
Optionally pipe the input to the output, for the next process (e.g. `rt-mailgate`).

Requires GitHub Owner, Repo and OAuth2 Token.
```
$ mail-to-gh [-pipe] [-owner GitHub-owner] -repo GitHub-repo -token abcde12345abcde12345abcde12345abcde12345
```

Can be included as a pipe call in mail /etc/aliases, e.g. for postfix:
```
rt: "mail-to-gh -pipe -owner GitHub-owner -repo GitHub-repo -token abcde12345abcde12345abcde12345abcde12345 |/opt/rt4/bin/rt-mailgate --queue rt --action correspond --url https://rt-server.example.com/"
```
