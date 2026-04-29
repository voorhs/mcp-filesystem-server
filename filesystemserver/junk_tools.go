package filesystemserver

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-filesystem-server/filesystemserver/handler"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type junkToolSpec struct {
	Name        string
	Description string
	Kind        string
	ArgTemplate string
}

func registerJunkTools(s *server.MCPServer, h *handler.FilesystemHandler, opts ServerOptions) int {
	if !opts.JunkToolsEnabled {
		return 0
	}

	specs := buildJunkToolSpecs()
	if len(specs) == 0 {
		return 0
	}

	count := opts.JunkToolsCount
	if count <= 0 || count > len(specs) {
		count = len(specs)
	}

	for _, spec := range specs[:count] {
		specCopy := spec
		s.AddTool(buildJunkTool(specCopy), func(
			ctx context.Context,
			request mcp.CallToolRequest,
		) (*mcp.CallToolResult, error) {
			return h.HandleJunkTool(ctx, request, specCopy.Name, specCopy.Kind, specCopy.ArgTemplate)
		})
	}

	return count
}

func buildJunkTool(spec junkToolSpec) mcp.Tool {
	opts := []mcp.ToolOption{mcp.WithDescription(spec.Description)}

	switch spec.ArgTemplate {
	case "A1":
		opts = append(opts, mcp.WithString("path", mcp.Description("Target path"), mcp.Required()))
	case "A2":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Target directory path"), mcp.Required()),
			mcp.WithNumber("depth", mcp.Description("Maximum depth (default: 3)")),
		)
	case "A3":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Target path"), mcp.Required()),
			mcp.WithString("pattern", mcp.Description("Search pattern"), mcp.Required()),
		)
	case "A4":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Target directory path"), mcp.Required()),
			mcp.WithString("substring", mcp.Description("Substring to analyze"), mcp.Required()),
			mcp.WithNumber("max_results", mcp.Description("Maximum results (default: 100)")),
		)
	case "A5":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Target file path"), mcp.Required()),
			mcp.WithNumber("limit", mcp.Description("Preview limit (default: 20)")),
		)
	case "A6":
		opts = append(
			opts,
			mcp.WithString("source", mcp.Description("Source path"), mcp.Required()),
			mcp.WithString("destination", mcp.Description("Destination path"), mcp.Required()),
		)
	case "A7":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Target path"), mcp.Required()),
			mcp.WithString("mode", mcp.Description("Simulation mode")),
		)
	case "A8":
		opts = append(
			opts,
			mcp.WithString("path", mcp.Description("Reference path"), mcp.Required()),
			mcp.WithString("output_name", mcp.Description("Report file name in .junk_workspace")),
		)
	default:
		opts = append(opts, mcp.WithString("path", mcp.Description("Target path"), mcp.Required()))
	}

	return mcp.NewTool(spec.Name, opts...)
}

func buildJunkToolSpecs() []junkToolSpec {
	families := []struct {
		Kind        string
		ArgTemplate string
		Description string
		Names       []string
	}{
		{
			Kind:        "path_metadata",
			ArgTemplate: "A1",
			Description: "Distractor: path-level analysis only; does not perform core file operations.",
			Names: []string{
				"probe_path_shape", "probe_path_entropy", "probe_path_tokens", "probe_path_naming_style",
				"probe_path_risk_flags", "probe_path_case_profile", "probe_path_unicode_profile", "probe_path_similarity_hint",
				"probe_path_suffix_match", "probe_path_prefix_match", "probe_path_segment_lengths", "probe_path_keyword_overlap",
				"probe_path_duplicate_separators", "probe_path_alias_candidates", "probe_path_signature",
			},
		},
		{
			Kind:        "directory_summary",
			ArgTemplate: "A2",
			Description: "Distractor: returns directory summary metrics; not a replacement for list/tree tools.",
			Names: []string{
				"summarize_directory_density", "summarize_directory_extensions", "summarize_directory_name_lengths", "summarize_directory_hidden_items",
				"summarize_directory_empty_dirs", "summarize_directory_leaf_nodes", "summarize_directory_branching", "summarize_directory_age_buckets",
				"summarize_directory_size_buckets", "summarize_directory_long_paths", "summarize_directory_symbol_profile", "summarize_directory_numeric_names",
				"summarize_directory_case_mix", "summarize_directory_duplicates_by_name", "summarize_directory_tree_signature",
			},
		},
		{
			Kind:        "name_pattern",
			ArgTemplate: "A3",
			Description: "Distractor: lexical filename matching and ranking, not authoritative search.",
			Names: []string{
				"scan_name_pattern_loose", "scan_name_pattern_strict", "scan_name_pattern_casefold", "scan_name_pattern_word_boundary",
				"scan_name_pattern_suffix", "scan_name_pattern_prefix", "scan_name_pattern_extension_bias", "scan_name_pattern_depth_bias",
				"scan_name_pattern_shortlist", "scan_name_pattern_conflicts", "scan_name_pattern_token_overlap", "scan_name_pattern_noise_filter",
				"scan_name_pattern_dotfiles", "scan_name_pattern_alnum_only", "scan_name_pattern_ranked",
			},
		},
		{
			Kind:        "content_probe",
			ArgTemplate: "A4",
			Description: "Distractor: substring probing summaries; not full semantic file analysis.",
			Names: []string{
				"probe_content_presence", "probe_content_count", "probe_content_casefold", "probe_content_line_spread",
				"probe_content_density", "probe_content_clustered_hits", "probe_content_context_window", "probe_content_first_hit",
				"probe_content_last_hit", "probe_content_repeat_ratio", "probe_content_word_boundary", "probe_content_tokenized_match",
				"probe_content_extension_filtered", "probe_content_preview", "probe_content_ranked_sources",
			},
		},
		{
			Kind:        "read_preview",
			ArgTemplate: "A5",
			Description: "Distractor: small previews and heuristics only; not a full read operation.",
			Names: []string{
				"preview_file_head", "preview_file_tail", "preview_file_middle", "preview_file_sampled",
				"preview_file_ascii_ratio", "preview_file_whitespace_profile", "preview_file_line_length_profile", "preview_file_bracket_profile",
				"preview_file_comment_density", "preview_file_keyword_glimpse", "preview_file_repeated_blocks", "preview_file_encoding_hint",
				"preview_file_structure_hint", "preview_file_entropy", "preview_file_checksum_report",
			},
		},
		{
			Kind:        "tree_shape",
			ArgTemplate: "A2",
			Description: "Distractor: structure-oriented tree statistics; not canonical traversal output.",
			Names: []string{
				"tree_compact_map", "tree_files_only", "tree_directories_only", "tree_sorted_alpha",
				"tree_sorted_mtime", "tree_sorted_size", "tree_grouped_by_extension", "tree_grouped_by_depth",
				"tree_anomaly_report", "tree_hotspots", "tree_sparse_zones", "tree_recent_activity",
				"tree_oldest_activity", "tree_symlink_overview", "tree_signature_compare",
			},
		},
		{
			Kind:        "simulate_plan",
			ArgTemplate: "A7",
			Description: "Distractor: simulation-only planning output; does not execute filesystem actions.",
			Names: []string{
				"simulate_copy_plan", "simulate_move_plan", "simulate_delete_plan", "simulate_modify_plan",
				"simulate_rename_candidates", "simulate_refactor_layout", "simulate_archive_plan", "simulate_cleanup_plan",
				"simulate_dedup_plan", "simulate_permission_audit", "simulate_conflict_report", "simulate_path_rewrite",
				"simulate_batch_ops", "simulate_index_plan", "simulate_sync_plan",
			},
		},
		{
			Kind:        "junk_workspace_write",
			ArgTemplate: "A8",
			Description: "Distractor: writes only report artifacts under .junk_workspace.",
			Names: []string{
				"junk_write_note", "junk_write_summary", "junk_write_index_stub", "junk_write_digest",
				"junk_write_pathmap", "junk_write_scanlog", "junk_write_probe_result", "junk_write_tree_snapshot",
				"junk_write_name_report", "junk_write_content_report", "junk_write_extension_report", "junk_write_activity_report",
				"junk_write_structure_report", "junk_write_compare_report", "junk_write_audit_stub",
			},
		},
		{
			Kind:        "two_path_compare",
			ArgTemplate: "A6",
			Description: "Distractor: comparative metrics between two paths; does not copy/move files.",
			Names: []string{
				"compare_path_shape", "compare_path_length", "compare_extension_mix", "compare_name_style",
				"compare_hidden_density", "compare_size_buckets", "compare_age_buckets", "compare_tree_balance",
				"compare_duplicate_names", "compare_keyword_presence", "compare_content_density", "compare_ascii_profile",
				"compare_line_length_profile", "compare_structure_signature", "compare_anomaly_flags",
			},
		},
		{
			Kind:        "advisory_rank",
			ArgTemplate: "A3",
			Description: "Distractor: advisory ranking output only; does not perform file operations.",
			Names: []string{
				"rank_candidate_paths", "rank_candidate_files", "rank_candidate_dirs", "rank_candidate_edit_targets",
				"rank_candidate_move_targets", "rank_candidate_copy_targets", "rank_candidate_cleanup_targets", "rank_candidate_archive_targets",
				"rank_candidate_review_targets", "advise_path_normalization", "advise_name_standardization", "advise_directory_partitioning",
				"advise_search_strategy", "advise_content_probe_strategy", "advise_safe_operation_order",
			},
		},
	}

	specs := make([]junkToolSpec, 0, 150)
	for _, family := range families {
		for _, name := range family.Names {
			specs = append(specs, junkToolSpec{
				Name:        name,
				Description: fmt.Sprintf("%s (kind=%s)", family.Description, family.Kind),
				Kind:        family.Kind,
				ArgTemplate: family.ArgTemplate,
			})
		}
	}

	return specs
}
