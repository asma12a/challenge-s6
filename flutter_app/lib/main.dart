import 'package:dio/dio.dart';
import 'package:dio_cache_interceptor/dio_cache_interceptor.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart' as provider;
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/screens/sign_in.dart';
import 'package:squad_go/screens/tabs.dart';
import 'package:google_fonts/google_fonts.dart';

final log = Logger("AppLogger");
final dio = Dio(BaseOptions(
  connectTimeout: Duration(seconds: 5),
  receiveTimeout: Duration(seconds: 5),
  headers: {
    'Accept': 'application/json',
  },
));

void main() async {
  await dotenv.load(fileName: "assets/../.env");
  dio.interceptors.add(
    DioCacheInterceptor(
      options: CacheOptions(
        store: MemCacheStore(),
        policy: CachePolicy.request,
      ),
    ),
  );
  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return provider.MultiProvider(
      providers: [
        provider.ChangeNotifierProvider(create: (_) => AuthState()),
      ],
      builder: (context, child) => MaterialApp(
        theme: theme,
        home: provider.Consumer<AuthState>(builder: (context, authState, _) {
          return authState.isAuthenticated
              ? const TabsScreen()
              : FutureBuilder(
                  future: authState.tryLogin(),
                  builder: (context, snapshot) =>
                      snapshot.connectionState == ConnectionState.waiting
                          ? const Scaffold(
                              body: Center(
                                child: CircularProgressIndicator(),
                              ),
                            )
                          : const SignInScreen(),
                );
        }),
        debugShowCheckedModeBanner: false,
      ),
    );
  }
}

final theme = ThemeData(
  useMaterial3: true,
  colorScheme: ColorScheme.fromSeed(
    seedColor: const Color.fromRGBO(8, 95, 113, 1),
  ),
  textTheme: GoogleFonts.latoTextTheme(),
);
