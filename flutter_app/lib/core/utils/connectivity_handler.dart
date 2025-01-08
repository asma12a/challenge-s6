import 'dart:async';
import 'dart:io';

import 'package:connectivity_plus/connectivity_plus.dart';
import 'package:flutter/foundation.dart';

class ConnectivityHandler {
  static final _singleton = ConnectivityHandler._internal();

  StreamSubscription<List<ConnectivityResult>>? _subscription;

  StreamController<bool> connectionChangeStatusController = StreamController();

  final Connectivity _connectivity = Connectivity();

  ConnectivityHandler._internal();

  bool _isConnected = false;

  bool get isConnected => _isConnected;

  factory ConnectivityHandler() {
    return _singleton;
  }

  dispose() {
    _subscription?.cancel();
    connectionChangeStatusController.close();
  }

  void initialize() async {
    assert(_subscription == null, 'Already connectivity handler initialized');

    final connectivityResult = await _connectivity.checkConnectivity();
    if (!connectivityResult.contains(ConnectivityResult.none)) {
      _isConnected = true;
    }
    _connectivityListener();
  }

  void _connectivityListener() {
    _subscription = _connectivity.onConnectivityChanged
        .listen((List<ConnectivityResult> results) {
      late bool currentState;
      if (results.contains(ConnectivityResult.none)) {
        currentState = false;
      } else {
        currentState = true;
      }

      if (currentState != _isConnected) {
        _isConnected = currentState;
        connectionChangeStatusController.add(_isConnected);
      }
    });
  }
}

class DeviceConfiguration {
  static bool get isOfflineSupportedDevice {
    if (kIsWeb) return false;
    return Platform.isIOS || Platform.isMacOS || Platform.isAndroid;
  }
}
