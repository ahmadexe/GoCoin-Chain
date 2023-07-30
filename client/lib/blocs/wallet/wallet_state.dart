part of 'wallet_bloc.dart';

@immutable
abstract class WalletState {
  final Wallet? wallet;
  final String? message;

  const WalletState({
    this.wallet,
    this.message,
  });
}

class WalletInitial extends WalletState {}

class WalletLoading extends WalletState {}

class WalletLoaded extends WalletState {
  const WalletLoaded({
    required Wallet wallet,
  }) : super(wallet: wallet);
}

class WalletError extends WalletState {
  const WalletError({
    required String message,
  }) : super(message: message);
}
