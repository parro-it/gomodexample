# gomod

This is a template I use to create my go module templates.

# Usage 

TLDR: click on 'use template' button above, have fun!

# How it work

When you create a new repository from this template,
[initialize](.github/workflows/initialize.yml) GitHub action 
is executed on your newly created repository.

The `initialize` action install and run 
[templatedir](github.com/parro-it/templatedir)
command (you can look there for details of what it 
does, but it basically render any file 
in the repository which has a `.template` 
extension and a Go template syntax. 

Rendered templates are saved to disk 
with the same filename but without `.template`.
Template files are deleted, and then all these changes
are committed back to your repository.

# Customize this template.

You can have a look at the `.template`
files in this repository as examples
of use if you want to make your own template
based on [templatedir](github.com/parro-it/templatedir).


# License
[MIT Licensed](LICENSE)

© 2021 Andrea Parodi



