

#[derive(Debug)]
pub struct ConsumerConfig {
    pub broker: String,
    pub group_id: String,
    pub topic: String,
}

impl ConsumerConfig {
    pub fn load() -> Self {
       Self { 
            broker: "host.docker.internal:9094".to_string(),
            group_id: "my_group".to_string(),
            topic: "nodeTopic".to_string(),
        }
    }
}