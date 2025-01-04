import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/platform/mobile/main_mobile.dart';
import 'package:squad_go/platform/web/main_web.dart';

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
