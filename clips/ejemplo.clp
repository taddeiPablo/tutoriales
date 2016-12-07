(deffunction area-sphere
        (?radius)
        (bind ?area (* (* (pi) 2)
                (* ?radius ?radius)))
        (return ?area))