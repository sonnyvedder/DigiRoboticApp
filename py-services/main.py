from fastapi import FastAPI
from pydantic import BaseModel
import uvicorn
import os
from dotenv import load_dotenv

# Load environment variables
load_dotenv()

app = FastAPI()

class VisionData(BaseModel):
    robot_id: str
    image_data: str

@app.get("/health")
async def health_check():
    return {"status": "ok"}

@app.post("/vision/process")
async def process_vision(data: VisionData):
    return {
        "status": "processed",
        "robot_id": data.robot_id,
        "detected_objects": []
    }

if __name__ == "__main__":
    port = int(os.getenv("PORT", 8081))
    uvicorn.run(app, host="0.0.0.0", port=port)