import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart' as provider;
import 'package:squad_go/core/models/event.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/platform/mobile/screens/event.dart';
import 'package:squad_go/platform/mobile/screens/join.dart';
import 'package:squad_go/shared_widgets/sign_in.dart';
import 'package:squad_go/platform/mobile/screens/tabs.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

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
          path: 'home',
          builder: (BuildContext context, GoRouterState state) {
            return const TabsScreen(
              initialPageIndex: 0,
            );
          },
        ),
        GoRoute(
          path: 'search',
          builder: (BuildContext context, GoRouterState state) {
            return const TabsScreen(
              initialPageIndex: 1,
            );
          },
        ),
        GoRoute(
          path: 'join',
          builder: (BuildContext context, GoRouterState state) {
            return const JoinEventScreen();
          },
        ),
        GoRoute(
          path: 'event/:id',
          builder: (BuildContext context, GoRouterState state) {
            final eventId = state.pathParameters['id'];
            final eventData = state.extra as Event?;
            return EventScreen(
              eventId: eventId!,
              event: eventData,
            );
          },
        ),
      ],
    ),
    GoRoute(
      path: '/sign-in',
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
        localizationsDelegates: AppLocalizations.localizationsDelegates,
        supportedLocales: AppLocalizations.supportedLocales,
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
