part of '../wallet_bloc.dart';

@immutable
abstract class WalletInfoState {
  final Wallet? wallet;
  final String? message;

  const WalletInfoState({
    this.wallet,
    this.message,
  });
}

class WalletInfoDefault extends WalletInfoState {
  const WalletInfoDefault() : super();
}

class WalletInfoLoading extends WalletInfoState {}

class WalletInfoLoaded extends WalletInfoState {
  const WalletInfoLoaded({
    required Wallet wallet,
  }) : super(wallet: wallet);
}

class WalletInfoError extends WalletInfoState {
  const WalletInfoError({
    required String message,
  }) : super(message: message);
}
