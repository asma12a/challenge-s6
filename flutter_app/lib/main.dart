import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/platform/mobile/screens/home.dart';
import 'package:squad_go/platform/web/screens/home.dart';

void main() async {
    await dotenv.load(fileName: "assets/../.env");

  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => AuthState()),
      ],
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    // Détecte la plateforme (web ou mobile)
    if (kIsWeb) {
      return const MyAppWeb();
    } else {
      return const MyAppMobile();
    }
  }
}

// Application pour Mobile
class MyAppMobile extends StatelessWidget {
  const MyAppMobile({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Application Mobile',
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
      ),
      home: const HomeScreen(),
      debugShowCheckedModeBanner: false,
    );
  }
}

// Application pour Web
class MyAppWeb extends StatelessWidget {
  const MyAppWeb({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Application Web',
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.teal),
      ),
      home: const WebHomeScreen(),
      debugShowCheckedModeBanner: false,
    );
  }
}
