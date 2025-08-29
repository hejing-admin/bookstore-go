quick start
```go
	options := &LoggerOptions{
		Level:      "debug",
		LogsDir:    "logs",
		MaxSize:    10,
		MaxAge:     5,
		MaxBackups: 5,
		Compress:   false,
		Console:    true,
	}

	mLogger, err := mlog.New(options)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	if mLogger.DebugEnable(){
		mLogger.L().Debug("debug message1")	
	}
	mLogger.L().Info("info message1")
	mLogger.L().Error("error message1")

	//update logger level
	mLogger.setLevel("info")
	mLogger.L().Debug("debug message2")
	mLogger.L().Info("info message2")
	mLogger.L().Error("error message2")
```

global function usage:
```go
    if err := mlog.Init(mlog.DefaultOptions()); err != nil {
        fmt.Printf("init mlog failed:%v\n", err)
    }
    
    // print raw string
    mlog.Debug("debug message 1")
    mlog.Info("info message 1")
    mlog.Warn("warn message 1")
    mlog.Error("error message 1")
    
    // print fmt string
    if mlog.DebugEnable() {
        mlog.Debugf("debug message %d", 2)
    }
    if mlog.InfoEnable() {
        mlog.Infof("info message %d", 2)
    }
    mlog.Warnf("warn message %d", 2)
    mlog.Errorf("error message %d", 2)
```
