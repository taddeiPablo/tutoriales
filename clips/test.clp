(deftemplate mensaje
        (slot nombre)
        (slot respuesta))

(deffacts mensajes
        (mensaje (nombre "hola") (respuesta "bien en que lo puedo ayudar"))
        (mensaje (nombre "buenos dias") (respuesta "buenos dias que necesita")))

(defrule dialogo
        (mensaje (nombre ?nom))
        => 
        (printout t ?nom crlf))