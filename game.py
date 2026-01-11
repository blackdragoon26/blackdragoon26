import os
import json

def handle_move(move):
    with open('state.json', 'r') as f:
        data = json.load(f)

    if move == "move_up":
        data['y'] -= 1
    
    # Logic to redraw the SVG based on new X,Y coordinates
    generate_svg(data['x'], data['y'])
    
    with open('state.json', 'w') as f:
        json.dump(data, f)

if __name__ == "__main__":
    # Get the move from the Issue Title environment variable
    move_title = os.getenv('ISSUE_TITLE') 
    handle_move(move_title)
