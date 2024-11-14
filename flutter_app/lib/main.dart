import 'package:flutter/material.dart';
import 'package:flutter_app/screens/sign_in.dart';
import 'package:flutter_app/screens/tabs.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

final theme = ThemeData(
  useMaterial3: true,
  colorScheme: ColorScheme.fromSeed(
    seedColor: const Color.fromARGB(255, 8, 95, 113),
  ),
  textTheme: GoogleFonts.latoTextTheme(),
);

Future main() async {
  await dotenv.load(fileName: "assets/../.env");
  runApp(
    const ProviderScope(
      child: App(),
    ),
  );
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: theme,
      home: const TabsScreen(),
      debugShowCheckedModeBanner: false,
    );
  }
}
