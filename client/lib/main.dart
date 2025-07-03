import 'package:dio/dio.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
      ),
      home: const MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
  TextEditingController controller = TextEditingController();

  List allValues = [];

  Map currentValues = {};

  final dio = Dio(BaseOptions(baseUrl: "http://localhost:3050/"));

  @override
  void initState() {
    super.initState();
    fetchAllValues();
    fetchCurrentValues();
  }

  void fetchCurrentValues() {
    dio.get("api/values/current").then((e) {
      setState(() {
        var data = e.data["values"];
        currentValues = switch (data) {
          Map() => data,
          _ => {},
        };
      });
    });
  }

  void fetchAllValues() {
    dio.get("api/values/all").then((e) {
      setState(() {
        var data = e.data["values"];
        allValues = switch (data) {
          List() => data,
          _ => <int>[],
        };
      });
    });
  }

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,

        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text('You have pushed the button this many times:'),
            Text("all values: $allValues"),
            Text("current values: $currentValues"),
            SizedBox(
              width: 200,
              child: TextField(
                controller: controller,
                onSubmitted: (_) => handle(),
                autofocus: true,
              ),
            ),
            ElevatedButton(onPressed: handle, child: Text("Submit")),
            Text(
              '$_counter',
              style: Theme.of(context).textTheme.headlineMedium,
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.add),
      ),
    );
  }

  void handle() {
    dio.post("api/values", data: {"value": int.tryParse(controller.text)}).then(
      (e) {
        fetchAllValues();
        fetchCurrentValues();
        controller.clear();
      },
    );
  }
}
