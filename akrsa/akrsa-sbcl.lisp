
(defun akrsa ()
    (let* ((bitsize (expt 2 4096))
      (e (random bitsize))
      (d (random bitsize))
      (m (random bitsize))
      (x (random bitsize))
      (r (mod (* e d) m))
      (c (mod (* e x) m))
      (y1 (mod (* c d) m))
      (y2 (mod (* r x) m)))
  (format t "~d~%" y1)
  (format t "~d~%" y2)
  (format t "~d~%" (= y1 y2))))

(sb-ext:save-lisp-and-die 
  "akrsa-lisp"
  ;; this is the main function:
  :toplevel (lambda () (akrsa))
  :executable t :purify t)
