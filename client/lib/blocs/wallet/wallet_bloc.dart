import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'wallet_event.dart';
part 'wallet_state.dart';

part 'data_provider.dart';

class WalletBloc extends Bloc<WalletEvent, WalletState> {
  WalletBloc() : super(WalletInitial()) {
    on<WalletEvent>((event, emit) {
      // TODO: implement event handler
    });
  }
}
