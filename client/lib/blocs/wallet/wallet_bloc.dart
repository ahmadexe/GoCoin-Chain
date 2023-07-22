import 'package:client/models/wallet.dart';
import 'package:dio/dio.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'wallet_event.dart';
part 'wallet_state.dart';

part 'data_provider.dart';

class WalletBloc extends Bloc<WalletEvent, WalletState> {
  static WalletBloc get(context, [listen = true]) =>
      BlocProvider.of(context, listen: listen);

  WalletBloc() : super(WalletInitial()) {
    on<WalletEvent>((event, emit) {});
  }

  Future<void> getWalletDetails(
      GetWalletDetails event, Emitter<WalletState> emit) async {
    emit(
      WalletLoading(),
    );

    try {
      final wallet = await _WallterProvider.getWalletDetails();

      emit(
        WalletLoaded(wallet: wallet),
      );
    } catch (e) {
      debugPrint('----- ERROR in Wallet Details Provider -----');
      debugPrint(e.toString());
      emit(
        const WalletError(message: 'Failed to get wallet details'),
      );
    }
  }
}
