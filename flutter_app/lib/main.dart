import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart' as provider;
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/screens/home.dart';
import 'package:squad_go/screens/join.dart';
import 'package:squad_go/screens/sign_in.dart';
import 'package:squad_go/screens/tabs.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:go_router/go_router.dart';

void main() async {
  await dotenv.load(fileName: "assets/../.env");
  runApp(const App());
}

/// The route configuration.
final GoRouter _router = GoRouter(
  redirect: (BuildContext context, GoRouterState state) async {
    final authState = provider.Provider.of<AuthState>(context, listen: false);

    if (authState.isAuthenticated || state.uri.path == '/sign-in') {
      return null;
    }

    final isLoggedIn = await authState.tryLogin();

    if (isLoggedIn) {
      return null;
    }
    return '/sign-in';
  },
  routes: <RouteBase>[
    GoRoute(
      path: '/',
      builder: (BuildContext context, GoRouterState state) {
        return const TabsScreen(); // Page principale après connexion
      },
      routes: <RouteBase>[
        GoRoute(
          path: 'join',
          builder: (BuildContext context, GoRouterState state) {
            return const JoinEventScreen();
          },
        ),
        GoRoute(
          path: 'tabs',
          builder: (BuildContext context, GoRouterState state) {
            return const TabsScreen();
          },
        ),
        GoRoute(
          path: 'tabs',
          builder: (BuildContext context, GoRouterState state) {
            return const TabsScreen();
          },
        ),
        GoRoute(
          path: 'event/:id',
          builder: (BuildContext context, GoRouterState state) {
            return const TabsScreen();
          },
        ),
        GoRoute(
          path: 'sign-in',
          redirect: (BuildContext context, GoRouterState state) {
            final authState =
                provider.Provider.of<AuthState>(context, listen: false);

            // Redirection vers /tabs si déjà connecté
            if (authState.isAuthenticated) {
              return '/';
            }
            return null;
          },
          builder: (BuildContext context, GoRouterState state) {
            return const SignInScreen();
          },
        ),
      ],
    ),
  ],
);

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return provider.MultiProvider(
      providers: [
        provider.ChangeNotifierProvider(create: (_) => AuthState()),
      ],
      builder: (context, child) => MaterialApp.router(
        routerConfig: _router,
        theme: theme,
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
