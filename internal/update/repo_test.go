package update

import (
	"os"
	"testing"
)

// The fork's updater must default to the fork. An upstream rebase that
// resolves a conflict in favor of asheshgoplani/agent-deck fails here.
func TestGitHubRepoDefaultsToFork(t *testing.T) {
	if DefaultGitHubRepo != "nirokato/agent-deck" {
		t.Fatalf("DefaultGitHubRepo = %q, want nirokato/agent-deck", DefaultGitHubRepo)
	}
	if _, set := os.LookupEnv("AGENTDECK_UPDATE_REPO"); !set && GitHubRepo != DefaultGitHubRepo {
		t.Fatalf("GitHubRepo = %q with AGENTDECK_UPDATE_REPO unset, want %q", GitHubRepo, DefaultGitHubRepo)
	}
}

func TestUpdateRepoEnvOverride(t *testing.T) {
	t.Setenv("AGENTDECK_UPDATE_REPO", " asheshgoplani/agent-deck ")
	if got := envOr("AGENTDECK_UPDATE_REPO", DefaultGitHubRepo); got != "asheshgoplani/agent-deck" {
		t.Fatalf("envOr = %q, want asheshgoplani/agent-deck", got)
	}
}
