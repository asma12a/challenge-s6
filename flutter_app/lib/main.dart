import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/platform/mobile/screens/home.dart';
import 'package:squad_go/platform/web/screens/home.dart';
import 'package:squad_go/shared_widgets/sign_in.dart'; 

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
    if (kIsWeb) {
      return const MyAppWeb();
    } else {
      return const App();
    }
  }
}

// Application mobile qui gère l'authentification
class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        useMaterial3: true,
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
      ),
      home: Consumer<AuthState>(
        builder: (context, authState, _) {
          debugPrint('isAuthenticated? : ${authState.isAuthenticated}');
          debugPrint('isAdmin? : ${authState.isAdmin}');

          if (authState.isAuthenticated) {
            return const HomeScreen();
          } else {
            return FutureBuilder(
              future: authState.tryLogin(), 
              builder: (context, snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return const Scaffold(
                    body: Center(
                      child: CircularProgressIndicator(),
                    ),
                  );
                } else {
                  return const SignInScreen();
                }
              },
            );
          }
        },
      ),
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
