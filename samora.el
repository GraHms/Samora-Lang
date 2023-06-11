;;; samora.el --- mode for editing samora scripts

;; Copyright (C) 2023 Ismael GraHms

;; Author: Ismael GraHms <grahms@outlook.com>
;; Keywords: languages
;; Version: 1.0

;;; Commentary:

;; Provides support for editing samora scripts with full support for
;; font-locking, but no special keybindings, or indentation handling.

;;;; Enabling:

;; Add the following to your .emacs file

;; (require 'samora)
;; (setq auto-mode-alist (append '(("\\.mon$" . samora-mode)) auto-mode-alist)))



;;; Code:

(defvar samora-constants
  '("true"
    "false"))

(defvar samora-keywords
  '(
    "else"
    "fn"
    "for"
    "foreach"
    "function"
    "if"
    "in"
    "let"
    "return"
    ))

;; The language-core and functions from the standard-library.
(defvar samora-functions
  '(
    "args"
    "exit"
    "file.close"
    "file.lines"
    "file.open"
    "first"
    "int"
    "last"
    "len"
    "math.abs"
    "math.random"
    "math.sqrt"
    "push"
    "puts"
    "read"
    "rest"
    "set"
    "string"
    "string.interpolate"
    "string.reverse"
    "string.split"
    "string.tolower"
    "string.toupper"
    "string.trim"
    "type"
    "version"
    ))


(defvar samora-font-lock-defaults
  `((
     ("\"\\.\\*\\?" . font-lock-string-face)
     (";\\|,\\|=" . font-lock-keyword-face)
     ( ,(regexp-opt samora-keywords 'words) . font-lock-builtin-face)
     ( ,(regexp-opt samora-constants 'words) . font-lock-constant-face)
     ( ,(regexp-opt samora-functions 'words) . font-lock-function-name-face)
     )))

(define-derived-mode samora-mode fundamental-mode "samora script"
  "samora-mode is a major mode for editing samora scripts"
  (setq font-lock-defaults samora-font-lock-defaults)

  ;; Comment handler for single & multi-line modes
  (modify-syntax-entry ?\/ ". 124b" samora-mode-syntax-table)
  (modify-syntax-entry ?\* ". 23n" samora-mode-syntax-table)

  ;; Comment ender for single-line comments.
  (modify-syntax-entry ?\n "> b" samora-mode-syntax-table)
  (modify-syntax-entry ?\r "> b" samora-mode-syntax-table)
  )

(provide 'samora)