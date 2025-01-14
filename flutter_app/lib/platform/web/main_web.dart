import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:squad_go/core/providers/auth_state_provider.dart';
import 'package:squad_go/core/services/auth_service.dart';
import 'package:squad_go/platform/web/screens/home.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';

class MyAppWeb extends StatelessWidget {
  const MyAppWeb({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => AuthState()),
        Provider<AuthService>(create: (_) => AuthService()),
      ],
      builder: (context, child) => MaterialApp(
        localizationsDelegates: AppLocalizations.localizationsDelegates,
        supportedLocales: AppLocalizations.supportedLocales,
        title: 'Squad GO',
        theme: ThemeData(
          useMaterial3: true,
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.blue),
        ),
        home: const WebHomeScreen(),
        debugShowCheckedModeBanner: false,
      ),
    );
  }
}
