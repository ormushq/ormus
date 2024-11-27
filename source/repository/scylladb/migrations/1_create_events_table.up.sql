CREATE TABLE event (
                             id UUID PRIMARY KEY,
                             type TEXT,
                             name TEXT,
                             properties MAP<TEXT, TEXT>,
                             send_at TIMESTAMP,
                             received_at TIMESTAMP,
                             timestamp TIMESTAMP,
                             event TEXT,
                             write_key TEXT,
                             created_at TIMESTAMP,
                             updated_at TIMESTAMP
);
