import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/screens/sign_in.dart';
import 'package:squad_go/screens/tabs.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

final theme = ThemeData(
  useMaterial3: true,
  colorScheme: ColorScheme.fromSeed(
    seedColor: const Color.fromRGBO(8, 95, 113, 1),
  ),
  textTheme: GoogleFonts.latoTextTheme(),
);

Future main() async {
  final storage = const FlutterSecureStorage();
  await dotenv.load(fileName: "assets/../.env");
  final token = dotenv.env['JWT_STORAGE KEY'] != null
      ? await storage.read(key: dotenv.env['JWT_STORAGE KEY']!)
      : null;

  runApp(
    ProviderScope(
      child: App(
        token: token,
      ),
    ),
  );
}

class App extends StatelessWidget {
  final String? token;
  const App({super.key, this.token});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: theme,
      home: token != null ? const TabsScreen() : const SignInScreen(),
      debugShowCheckedModeBanner: false,
    );
  }
}
