# Auxiliary Tool Catalog (150 tools)

This file proposes an experiment-ready set of 150 auxiliary tools for the filesystem MCP server.

Goals:
- Keep tools in-domain and plausible.
- Avoid overlapping with core task-critical behavior.
- Keep behavior safe and deterministic.
- Make semantic confusion realistic (name/description/arg-level).

Safety contract for all tools below:
- They MUST obey existing allowed-directory validation.
- They MUST NOT modify files unless explicitly marked as `junk workspace write`.
- Any write-capable distractor MUST write only under `.junk_workspace/`.
- They MUST return clearly marked payloads, e.g. `{ "distractor": true, ... }`.

---

## Argument templates

To keep implementation simple, tools reuse these argument shapes:

- `A1`: `path` (required)
- `A2`: `path` (required), `depth` (optional, default: 3)
- `A3`: `path` (required), `pattern` (required)
- `A4`: `path` (required), `substring` (required), `max_results` (optional, default: 100)
- `A5`: `path` (required), `limit` (optional, default: 20)
- `A6`: `source` (required), `destination` (required)
- `A7`: `path` (required), `mode` (optional, enum-ish string)
- `A8`: `path` (required), `output_name` (optional)

---

## 1) Path/metadata utilities (1-15)

1. `probe_path_shape` - Args: `A1` - Return normalized path segments and depth classification.
2. `probe_path_entropy` - Args: `A1` - Return simple character entropy estimate for path string.
3. `probe_path_tokens` - Args: `A1` - Return tokenized path components and extension hints.
4. `probe_path_naming_style` - Args: `A1` - Infer snake/camel/kebab style from basename.
5. `probe_path_risk_flags` - Args: `A1` - Return heuristic warnings (hidden, long-name, temp-like).
6. `probe_path_case_profile` - Args: `A1` - Return uppercase/lowercase ratios for basename.
7. `probe_path_unicode_profile` - Args: `A1` - Return unicode category summary for basename.
8. `probe_path_similarity_hint` - Args: `A3` - Score similarity of basename against pattern text.
9. `probe_path_suffix_match` - Args: `A3` - Return whether basename suffix resembles pattern.
10. `probe_path_prefix_match` - Args: `A3` - Return whether basename prefix resembles pattern.
11. `probe_path_segment_lengths` - Args: `A1` - Return list of segment lengths for path.
12. `probe_path_keyword_overlap` - Args: `A3` - Return keyword overlap stats for basename vs pattern.
13. `probe_path_duplicate_separators` - Args: `A1` - Report repeated separator anomalies.
14. `probe_path_alias_candidates` - Args: `A1` - Suggest alias names without creating anything.
15. `probe_path_signature` - Args: `A1` - Return stable hash-like signature of path string only.

## 2) Directory summary utilities (16-30)

16. `summarize_directory_density` - Args: `A2` - Return file/dir count ratios by depth.
17. `summarize_directory_extensions` - Args: `A2` - Return extension frequency histogram.
18. `summarize_directory_name_lengths` - Args: `A2` - Return basename length distribution.
19. `summarize_directory_hidden_items` - Args: `A2` - Return count and sample of dotfiles.
20. `summarize_directory_empty_dirs` - Args: `A2` - Return empty-directory candidates.
21. `summarize_directory_leaf_nodes` - Args: `A2` - Return count of leaf files and leaf dirs.
22. `summarize_directory_branching` - Args: `A2` - Return branching-factor stats.
23. `summarize_directory_age_buckets` - Args: `A2` - Return mtime bucket distribution.
24. `summarize_directory_size_buckets` - Args: `A2` - Return size bucket distribution.
25. `summarize_directory_long_paths` - Args: `A2` - Return top longest paths.
26. `summarize_directory_symbol_profile` - Args: `A2` - Return punctuation/symbol usage in names.
27. `summarize_directory_numeric_names` - Args: `A2` - Return count of numerically named items.
28. `summarize_directory_case_mix` - Args: `A2` - Return mixed-case naming statistics.
29. `summarize_directory_duplicates_by_name` - Args: `A2` - Return duplicate basenames across branches.
30. `summarize_directory_tree_signature` - Args: `A2` - Return compact deterministic tree fingerprint.

## 3) Filename-pattern utilities (31-45)

31. `scan_name_pattern_loose` - Args: `A3` - Return loose fuzzy filename matches.
32. `scan_name_pattern_strict` - Args: `A3` - Return strict substring filename matches.
33. `scan_name_pattern_casefold` - Args: `A3` - Return case-insensitive filename matches.
34. `scan_name_pattern_word_boundary` - Args: `A3` - Return word-boundary-aware name matches.
35. `scan_name_pattern_suffix` - Args: `A3` - Return filename suffix matches.
36. `scan_name_pattern_prefix` - Args: `A3` - Return filename prefix matches.
37. `scan_name_pattern_extension_bias` - Args: `A3` - Return matches prioritized by extension.
38. `scan_name_pattern_depth_bias` - Args: `A3` - Return matches prioritized by shallow depth.
39. `scan_name_pattern_shortlist` - Args: `A5` - Return top-k probable name candidates.
40. `scan_name_pattern_conflicts` - Args: `A3` - Return near-collision filenames.
41. `scan_name_pattern_token_overlap` - Args: `A3` - Return token-overlap score per filename.
42. `scan_name_pattern_noise_filter` - Args: `A3` - Return matches with temp/cache folders deprioritized.
43. `scan_name_pattern_dotfiles` - Args: `A3` - Return dotfile-only pattern matches.
44. `scan_name_pattern_alnum_only` - Args: `A3` - Return alphanumeric-normalized matches.
45. `scan_name_pattern_ranked` - Args: `A3` - Return ranked filename candidates with simple reasons.

## 4) Content-probe utilities (46-60)

46. `probe_content_presence` - Args: `A4` - Return whether substring appears at least once.
47. `probe_content_count` - Args: `A4` - Return approximate substring occurrence counts.
48. `probe_content_casefold` - Args: `A4` - Return case-insensitive occurrence counts.
49. `probe_content_line_spread` - Args: `A4` - Return matched line number spread stats.
50. `probe_content_density` - Args: `A4` - Return match density per file.
51. `probe_content_clustered_hits` - Args: `A4` - Return whether hits are clustered vs sparse.
52. `probe_content_context_window` - Args: `A4` - Return tiny windows around top matches.
53. `probe_content_first_hit` - Args: `A4` - Return first-hit location per file.
54. `probe_content_last_hit` - Args: `A4` - Return last-hit location per file.
55. `probe_content_repeat_ratio` - Args: `A4` - Return repeated-line ratio around hits.
56. `probe_content_word_boundary` - Args: `A4` - Return boundary-aware match counts.
57. `probe_content_tokenized_match` - Args: `A4` - Return tokenized approximate matches.
58. `probe_content_extension_filtered` - Args: `A4` - Return matches with extension histogram.
59. `probe_content_preview` - Args: `A4` - Return short snippets for top matches only.
60. `probe_content_ranked_sources` - Args: `A4` - Return ranked candidate files likely containing substring.

## 5) Read-preview utilities (61-75)

61. `preview_file_head` - Args: `A5` - Return first N lines with truncation marker.
62. `preview_file_tail` - Args: `A5` - Return last N lines with truncation marker.
63. `preview_file_middle` - Args: `A5` - Return middle window lines only.
64. `preview_file_sampled` - Args: `A5` - Return evenly sampled lines from file.
65. `preview_file_ascii_ratio` - Args: `A1` - Return ascii/non-ascii ratio, not full content.
66. `preview_file_whitespace_profile` - Args: `A1` - Return whitespace and indentation statistics.
67. `preview_file_line_length_profile` - Args: `A1` - Return line length distribution summary.
68. `preview_file_bracket_profile` - Args: `A1` - Return bracket/paren usage counts.
69. `preview_file_comment_density` - Args: `A1` - Return comment-like line density estimate.
70. `preview_file_keyword_glimpse` - Args: `A3` - Return tiny snippets around keyword hits.
71. `preview_file_repeated_blocks` - Args: `A1` - Return repeated-line block hints.
72. `preview_file_encoding_hint` - Args: `A1` - Return guessed encoding and confidence.
73. `preview_file_structure_hint` - Args: `A1` - Return rough text structure markers.
74. `preview_file_entropy` - Args: `A1` - Return line-level entropy summary.
75. `preview_file_checksum_report` - Args: `A1` - Return checksum plus basic size metadata.

## 6) Tree-shape utilities (76-90)

76. `tree_compact_map` - Args: `A2` - Return compact tree with only names and depth.
77. `tree_files_only` - Args: `A2` - Return files-only tree skeleton.
78. `tree_directories_only` - Args: `A2` - Return directories-only tree skeleton.
79. `tree_sorted_alpha` - Args: `A2` - Return tree sorted alphabetically.
80. `tree_sorted_mtime` - Args: `A2` - Return tree sorted by modification time.
81. `tree_sorted_size` - Args: `A2` - Return tree sorted by item size estimate.
82. `tree_grouped_by_extension` - Args: `A2` - Return extension-grouped tree summary.
83. `tree_grouped_by_depth` - Args: `A2` - Return counts grouped by depth.
84. `tree_anomaly_report` - Args: `A2` - Return suspicious naming/path anomalies in tree.
85. `tree_hotspots` - Args: `A2` - Return subtrees with highest item density.
86. `tree_sparse_zones` - Args: `A2` - Return sparse/empty regions in tree.
87. `tree_recent_activity` - Args: `A2` - Return most recently touched paths.
88. `tree_oldest_activity` - Args: `A2` - Return oldest touched paths.
89. `tree_symlink_overview` - Args: `A2` - Return symlink count and locations summary.
90. `tree_signature_compare` - Args: `A6` - Return structural signature comparison of two directories.

## 7) Plan/simulation utilities (91-105)

91. `simulate_copy_plan` - Args: `A6` - Return hypothetical copy plan without copying.
92. `simulate_move_plan` - Args: `A6` - Return hypothetical move plan without moving.
93. `simulate_delete_plan` - Args: `A1` - Return hypothetical delete impact report.
94. `simulate_modify_plan` - Args: `A7` - Return hypothetical edit plan with no writes.
95. `simulate_rename_candidates` - Args: `A1` - Return candidate rename schemes.
96. `simulate_refactor_layout` - Args: `A1` - Return hypothetical directory layout improvements.
97. `simulate_archive_plan` - Args: `A1` - Return hypothetical archive groupings.
98. `simulate_cleanup_plan` - Args: `A1` - Return hypothetical cleanup suggestions.
99. `simulate_dedup_plan` - Args: `A1` - Return hypothetical duplicate cleanup actions.
100. `simulate_permission_audit` - Args: `A1` - Return hypothetical permission audit summary.
101. `simulate_conflict_report` - Args: `A6` - Return name/path conflict risks for source/destination.
102. `simulate_path_rewrite` - Args: `A7` - Return transformed path text only, no filesystem changes.
103. `simulate_batch_ops` - Args: `A1` - Return pseudo batch operation schedule.
104. `simulate_index_plan` - Args: `A1` - Return plan for indexing files, no index written.
105. `simulate_sync_plan` - Args: `A6` - Return hypothetical sync steps and risk markers.

## 8) Report-writing utilities (106-120)

These are intentionally write-capable but MUST write only under `.junk_workspace/`.

106. `write_note_report` - Args: `A8` - Write a metadata note file into `.junk_workspace/`.
107. `write_summary_report` - Args: `A8` - Write a short summary artifact into `.junk_workspace/`.
108. `write_index_stub` - Args: `A8` - Write an index manifest into `.junk_workspace/`.
109. `write_digest_report` - Args: `A8` - Write digest-like report into `.junk_workspace/`.
110. `write_pathmap_report` - Args: `A8` - Write path mapping report into `.junk_workspace/`.
111. `write_scanlog_report` - Args: `A8` - Write scan log artifact into `.junk_workspace/`.
112. `write_probe_result` - Args: `A8` - Write probe output artifact into `.junk_workspace/`.
113. `write_tree_snapshot` - Args: `A8` - Write compact tree snapshot into `.junk_workspace/`.
114. `write_name_report` - Args: `A8` - Write filename-style report into `.junk_workspace/`.
115. `write_content_report` - Args: `A8` - Write content-hit report into `.junk_workspace/`.
116. `write_extension_report` - Args: `A8` - Write extension histogram into `.junk_workspace/`.
117. `write_activity_report` - Args: `A8` - Write mtime activity report into `.junk_workspace/`.
118. `write_structure_report` - Args: `A8` - Write structure profile report into `.junk_workspace/`.
119. `write_compare_report` - Args: `A8` - Write source/destination comparison report.
120. `write_audit_stub` - Args: `A8` - Write an audit placeholder report.

## 9) Two-path comparison utilities (121-135)

121. `compare_path_shape` - Args: `A6` - Compare path token and segment structures.
122. `compare_path_length` - Args: `A6` - Compare path-length and depth metrics.
123. `compare_extension_mix` - Args: `A6` - Compare extension distributions.
124. `compare_name_style` - Args: `A6` - Compare naming-style profiles.
125. `compare_hidden_density` - Args: `A6` - Compare hidden-file densities.
126. `compare_size_buckets` - Args: `A6` - Compare size bucket histograms.
127. `compare_age_buckets` - Args: `A6` - Compare mtime age-bucket histograms.
128. `compare_tree_balance` - Args: `A6` - Compare branching balance metrics.
129. `compare_duplicate_names` - Args: `A6` - Compare duplicate-basename pressure.
130. `compare_keyword_presence` - Args: `A3` - Compare keyword prevalence in one path.
131. `compare_content_density` - Args: `A4` - Compare substring density style metrics.
132. `compare_ascii_profile` - Args: `A6` - Compare ascii/non-ascii profile estimates.
133. `compare_line_length_profile` - Args: `A6` - Compare line-length distributions.
134. `compare_structure_signature` - Args: `A6` - Compare deterministic structure signatures.
135. `compare_anomaly_flags` - Args: `A6` - Compare anomaly-flag sets.

## 10) Advisory/ranking utilities (136-150)

136. `rank_candidate_paths` - Args: `A3` - Rank candidate paths by lexical relevance.
137. `rank_candidate_files` - Args: `A3` - Rank files likely related to pattern intent.
138. `rank_candidate_dirs` - Args: `A3` - Rank directories likely related to pattern intent.
139. `rank_candidate_edit_targets` - Args: `A3` - Rank files that look editable for pattern.
140. `rank_candidate_move_targets` - Args: `A3` - Rank plausible move destinations.
141. `rank_candidate_copy_targets` - Args: `A3` - Rank plausible copy destinations.
142. `rank_candidate_cleanup_targets` - Args: `A1` - Rank paths that look cleanup-worthy.
143. `rank_candidate_archive_targets` - Args: `A1` - Rank paths suitable for archiving.
144. `rank_candidate_review_targets` - Args: `A1` - Rank paths needing manual inspection.
145. `advise_path_normalization` - Args: `A1` - Return normalization suggestions, no changes.
146. `advise_name_standardization` - Args: `A1` - Return naming standardization hints.
147. `advise_directory_partitioning` - Args: `A1` - Return directory partitioning hints.
148. `advise_search_strategy` - Args: `A3` - Return strategy hints for finding relevant files.
149. `advise_content_probe_strategy` - Args: `A4` - Return strategy hints for content probing.
150. `advise_safe_operation_order` - Args: `A1` - Return safe op ordering hints, no execution.

---

## Notes for experiment reporting

- All tools are intentionally non-essential for benchmark tasks.
- Most tools are analytical summaries or planning helpers.
- The 15 junk-write tools are sandboxed to `.junk_workspace/`.
- The set is designed to increase lexical and semantic overlap while avoiding environment breakage.
