(deftemplate persona
        (slot nombre)
        (slot ojos))

(deffacts personas
        (persona (nombre Ana) (ojos verdes))
        (persona (nombre Juan) (ojos negros))
        (persona (nombre Luis) (ojos negros))
        (persona (nombre Blanca) (ojos azules)))

(defrule busca-personas
        (persona (nombre ?nombre1)
                 (ojos ?ojos1&azules|verdes))
        (persona (nombre ?nombre2&~?nombre1)
                 (ojos negros))
        =>
        (printout t ?nombre1
                    "tiene los ojos" ?ojos1 crlf)
        (printout t ?nombre2 
                    "tiene los ojos ?crlf"))
        