# ESPN GO API
ESPN's Fantasy Football [site](http://www.espn.com/nfl/) offers an API that developers can query for
information about leagues that ESPN hosts.  It's great ESPN offers this API, but
for the developers who knew about and used the "v2" API, yes, your code is
broken. What this project is looking to accomplish, is to improve on previous
attempts to create an API by providing a code representation of the "v3" ESPN
API in the form of generated (because it keep changing) granular GOlang structs
and to create stable API abstractions on top of those GOlang structs so
developers building projects will have a consistent, common, and tested API to
rely on.

### Generate the V3 API
```bash
make generateAPI
```

### Try out the in tree code consuming the API
```bash
make run
```
