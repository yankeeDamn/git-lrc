# Rule 1: git-lrc build command

Build and install lrc locally with:

make build-local && lrc hooks install

When building the lrc command, always use the binary name "lrc"

A primary rule in programming is that the name must match function or meaning. If the name says one thing and the function or meaning is another - it must be treated as a major bug and treated with highest priority. Always fix naming bugs immediately when discovered.

When doing refactors or improvements - don't do fallback implementations unless explicitly asked for. Having fallbacks creates very confusing program behavior.

When you want to run a test - always run the most specific test you want to run. Don't run all or other irrelevant tests.

When making proposals - think of how you can accomplish goals in an incremental way rather than suggesting sweeping refactors.
