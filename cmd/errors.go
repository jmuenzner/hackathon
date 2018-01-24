package cmd

import "errors"

var (
	// ErrNoNatsAddrs indicates that we have no Nats addresses to connect to.
	// These should either be supplied eplicitly, or discovered via DNS lookup.
	ErrNoNatsAddrs = errors.New("No nats server addresses available to connect " +
		"to. You should either provide these explicitly via flags, or implicitly " +
		"via DNS discovery")

	// ErrNoTargetNodeAddr indicates that an address of the node to contract
	// wasn't provided.
	ErrNoTargetNodeAddr = errors.New("You must provide the address " +
		"of an existing cluster node")

	// ErrBootstrapTimeout indicates that we timed out while waiting for a response
	// from the node that we asked to bootstrap for us.
	//
	// NOTE: whether bootstrapping actually ended up succeeding or failing is
	// unknown at this point. This case likely requires human intervention.
	ErrBootstrapTimeout = errors.New("Timed out while waiting for the target " +
		"cluster node to respond to the bootstrap request")

	// ErrBootstrapUnknownHost indicates that the host of the target node was
	// not known.
	ErrBootstrapUnknownHost = errors.New("The bootstrap target cluster node's " +
		"host was unknown")

	// ErrBootstrapConnectionRefused indicates that the connection to the target
	// cluster node was refused
	ErrBootstrapConnectionRefused = errors.New("Connection refused to target " +
		"cluster node")

	// ErrBootstrapEOF indicates that the connection was closed unexpectedly.
	ErrBootstrapEOF = errors.New("Connection was closed unexpectedly")
)
