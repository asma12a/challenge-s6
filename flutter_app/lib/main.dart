import 'dart:io';

import 'package:dio/dio.dart';
import 'package:dio_cache_interceptor/dio_cache_interceptor.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:squad_go/core/utils/connectivity_handler.dart';
import 'package:squad_go/platform/mobile/main_mobile.dart';
import 'package:squad_go/platform/web/main_web.dart';
import 'package:flutter_map_tile_caching/flutter_map_tile_caching.dart';

final log = Logger("AppLogger");
final dio = Dio(BaseOptions(
  validateStatus: (status) {
    return status! < 500; // Permet de recevoir les réponses 4xx sans exception.
  },
  connectTimeout: Duration(seconds: 30),
  receiveTimeout: Duration(seconds: 30),
  headers: {
    'Accept': 'application/json',
  },
));

final initialCacheOptions = CacheOptions(
  store: MemCacheStore(),
  policy: CachePolicy.request,
  priority: CachePriority.high,
  maxStale: const Duration(hours: 1),
  hitCacheOnErrorExcept: [401, 403],
  keyBuilder: (req) => req.uri.toString(),
);

void main() async {
  dio.interceptors.add(DioCacheInterceptor(options: initialCacheOptions));

  // Flutter Maps Tile Caching
  WidgetsFlutterBinding.ensureInitialized();
  if (!kIsWeb) {
    await FMTCObjectBoxBackend().initialise();
    await FMTCStore('mapStore').manage.create();
  }
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  void initState() {
    super.initState();

    ConnectivityHandler().initialize();
  }

  @override
  void dispose() {
    ConnectivityHandler().dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    if (kIsWeb) {
      return const MyAppWeb();
    } else {
      return const MyAppMobile();
    }
  }
}
