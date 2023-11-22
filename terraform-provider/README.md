my go.work file:

```
go 1.19

use (
	.
	..\go-admin
)

replace github.com/sweeneyb/preferences-finder/go-admin => C:/Users/sweeneyb/projects/preferences-finder/go-admin
replace github.com/sweeneyb/preferences-finder/go-admin/pflib => C:/Users/sweeneyb/projects/preferences-finder/go-admin/pflib
```