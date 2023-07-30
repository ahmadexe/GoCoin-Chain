part of 'wallet_bloc.dart';

class WalletState extends Equatable {
  final WalletInfoState wallet;

  const WalletState({
    required this.wallet,
  });

  WalletState copyWith({
    WalletInfoState? wallet,
  }) {
    return WalletState(
      wallet: wallet ?? this.wallet,
    );
  }
  
  @override
  List<Object?> get props => [wallet];
}

class WalletDefault extends WalletState {
  const WalletDefault() : super(wallet: const WalletInfoDefault());
}
