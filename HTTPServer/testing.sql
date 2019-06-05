SELECT  
    distinct history.identifier as identifier,
    reqNoEdit.method    AS requestMethod,
    reqNoEdit.path      AS requestPath,
    reqNoEdit.proto     AS requestProto,
    reqNoEdit.host      AS requestHost,
    reqNoEdit.header    AS requestHeaders,
    reqNoEdit.param     AS requestParam,

    reqEdit.method      AS requestEditMethod,
    reqEdit.path        AS requestEditPath,
    reqEdit.proto       AS requestEditProto,
    reqEdit.host        AS requestEditHost,
    reqEdit.header      AS requestEditHeaders,
    reqEdit.param       AS requestEditParam,

    resA.header AS responseHeaders,
    body
FROM "histories" as history
LEFT JOIN(
    SELECT 
        req.identifier as identifier, method, path, proto, host, header, param
    FROM "requests" as req
    LEFT JOIN (
        SELECT
            identifier,
            "{data:{"||group_concat(name||':"'||value||'"',",")||"}}" AS param
        FROM
            "request_data"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS reqd ON
        req.identifier = reqd.identifier
    LEFT JOIN (
        SELECT
            identifier,
            "{header:{"||group_concat(name||':"'||value||'"',",")||"}" AS header
        FROM
            "request_headers"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS reqh ON
        req.identifier = reqh.identifier
    WHERE is_edit = 0
) AS reqNoEdit ON
    history.identifier = reqNoEdit.identifier
LEFT JOIN(
    SELECT 
        req.identifier as identifier, method, path, proto, host, header, param
    FROM "requests" as req
    LEFT JOIN (
        SELECT
            identifier,
            "{data:{"||group_concat(name||':"'||value||'"',",")||"}}" AS param
        FROM
            "request_data"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS reqd ON
        req.identifier = reqd.identifier
    LEFT JOIN (
        SELECT
            identifier,
            "{header:{"||group_concat(name||':"'||value||'"',",")||"}" AS header
        FROM
            "request_headers"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS reqh ON
        req.identifier = reqh.identifier
    WHERE is_edit = 1
) AS reqEdit ON
    history.identifier = reqEdit.identifier
LEFT JOIN(
    SELECT res.identifier,body,header
    FROM "responses" as res
    LEFT JOIN (
        SELECT
            identifier,
            "{header:{"||group_concat(name||':"'||value||'"',",")||"}" AS header
        FROM
            "response_headers"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS resh ON
        res.identifier = resh.identifier
    LEFT JOIN (
        SELECT
            identifier, body
        FROM
            "response_bodies"
        GROUP BY identifier
        ORDER BY id ASC
    ) AS resb ON
        res.identifier = resb.identifier
) AS resA ON
    history.identifier = resA.identifier
WHERE history.identifier = "33a5525032b273ef9441454263bd55ef49b9d7dd"