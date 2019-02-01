# log

provides level distinguish and error detail logging.

These are the different logging levels.

#### PanicLevel level
highest level of severity. Logs and then calls panic with the message passed to Debug, Info, ...

#### FatalLevel level
Logs and then calls `os.Exit(1)`. It will exit even if the logging level is set to Panic.

#### ErrorLevel level
Used for errors that should definitely be noted.
Commonly used for hooks to send errors to an error tracking service.

#### WarnLevel level
Non-critical entries that deserve eyes.

#### InfoLevel level
General operational entries about what's going on inside the application.

#### DebugLevel level
Usually only enabled when debugging. Very verbose logging.
