## Semicolons

Go's formal grammar uses semicolons to terminate statements.

but those semicolons do not appear in the source.

the lexer(词法分析器) uses a simple rule to insert semicolons automatically as it scans.

Go programs have semicolons only in places such as for loop causes.

One consequence(做为什么的结果，重要性) of the semicolon insertion rules is that

you cannot put the open brace of a control structure on the next line.



