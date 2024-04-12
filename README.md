# Tidyshelf

## Roadmap

1. Implement an endpoint which accepts an image file and sends it to a configurable upstream service.
- TBD

2. The configurable upstream service could be for example Google's Gemini Pro 1.5. In early stages of development we will gladly use a mocked (and injected) dependency
- when sending a prompt to Gemini it is important to specify the format we want to get from themodel in return
- TBD

3. A next step could be coupling that web service with a LED matrix that highlights books on the bookshelf.
- use a colored tape to make it easier for Gemini's image recognition mechanism to distinguish different shelves from one another
- TBD
