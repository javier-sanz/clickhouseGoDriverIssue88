# clickhouseGoDriverIssue88
Code that tests issue https://github.com/kshvakov/clickhouse/issues/88

In order to reproduce: 

* Star clickhouse using docker as follows: `docker run -p9000:9000 yandex/clickhouse-server:1.1.54327`
* Execute the test code: `go run issue88Test.go`

The test will end up with a panic:

```
panic: Error committing: code: 101, message: Unexpected packet Hello received from client

goroutine 1 [running]:
main.main()
	/Users/jsanz/Projects/goWorkspace/src/clickhouseGoDriverIssue88/issue88Test.go:45 +0x2c4
exit status 2
```
And on Clickhouse log you will see the next stacktrace:

```
2018.02.10 09:25:02.027633 [ 3 ] <Information> TCPHandler: Processed in 0.001 sec.
2018.02.10 09:25:02.027926 [ 3 ] <Debug> executeQuery: (from 172.17.0.1:38078, query_id: 4e427f5e-ca52-4ac1-88a2-915a6f703214) INSERT INTO default.test  (EventDate, Value) VALUES
2018.02.10 09:25:02.040773 [ 3 ] <Error> executeQuery: Code: 101, e.displayText() = DB::Exception: Unexpected packet Hello received from client, e.what() = DB::Exception (from 172.17.0.1:38078) (in query: INSERT INTO default.test  (EventDate, Value) VALUES), Stack trace:

0. /usr/bin/clickhouse-server(StackTrace::StackTrace()+0x16) [0x3419e66]
1. /usr/bin/clickhouse-server(DB::Exception::Exception(std::__cxx11::basic_string<char, std::char_traits<char>, std::allocator<char> > const&, int)+0x1f) [0x14f2a9f]
2. /usr/bin/clickhouse-server(DB::TCPHandler::receivePacket()+0x146) [0x14fe146]
3. /usr/bin/clickhouse-server(DB::TCPHandler::readData(DB::Settings const&)+0x16a) [0x14fe52a]
4. /usr/bin/clickhouse-server(DB::TCPHandler::processInsertQuery(DB::Settings const&)+0x274) [0x14ff244]
5. /usr/bin/clickhouse-server(DB::TCPHandler::runImpl()+0x6ab) [0x14ffa7b]
6. /usr/bin/clickhouse-server(DB::TCPHandler::run()+0x2b) [0x150055b]
7. /usr/bin/clickhouse-server(Poco::Net::TCPServerConnection::start()+0xf) [0x3ae7c5f]
8. /usr/bin/clickhouse-server(Poco::Net::TCPServerDispatcher::run()+0xad) [0x3aecead]
9. /usr/bin/clickhouse-server(Poco::PooledThread::run()+0x86) [0x3a1bad6]
10. /usr/bin/clickhouse-server(Poco::ThreadImpl::runnableEntry(void*)+0x38) [0x39facf8]
11. /usr/bin/clickhouse-server() [0x4290b0f]
12. /lib/x86_64-linux-gnu/libpthread.so.0(+0x76ba) [0x7f906d3876ba]
13. /lib/x86_64-linux-gnu/libc.so.6(clone+0x6d) [0x7f906c9a83dd]
```

