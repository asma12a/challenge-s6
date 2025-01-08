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
  connectTimeout: Duration(seconds: 5),
  receiveTimeout: Duration(seconds: 5),
  headers: {
    'Accept': 'application/json',
  },
));

void main() async {
  ConnectivityHandler().initialize();

  dio.interceptors.add(
    DioCacheInterceptor(
      options: CacheOptions(
        store: MemCacheStore(),
        policy: CachePolicy.request,
      ),
    ),
  );

  // Flutter Maps Tile Caching
  WidgetsFlutterBinding.ensureInitialized();
  if (!kIsWeb) {
    await FMTCObjectBoxBackend().initialise();
    await FMTCStore('mapStore').manage.create();
  }
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    if (kIsWeb) {
      return const MyAppWeb();
    } else {
      return const MyAppMobile();
    }
  }
}
