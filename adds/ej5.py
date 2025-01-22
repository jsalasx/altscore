import requests

def parse_grid(encoded_message):
    """
    Decodifica el mensaje del radar y lo convierte en una matriz 8x8.
    """
    rows = encoded_message.split('|')
    rows.pop()  # Eliminar el último elemento vacío
    rows.reverse()  # Invertir filas para alinear con el sistema de coordenadas
    grid = []

    for index, row in enumerate(rows):
        # Dividir la fila en celdas de 3 caracteres
        new_row = split_into_chunks(row)
        # Tomar solo el segundo carácter de cada celda
        linea = [nRow[1] for nRow in new_row]
        grid.append(linea)

        # Mostrar la fila con colores (opcional, para depuración)
        texto = f"\033[34m{abs(index-8)}\033[0m "  # Número de fila
        for i, cell in enumerate(linea):
            if cell == '^':
                texto += f"\033[31m{cell}\033[0m "  # Enemigo en rojo
            elif cell == '#':
                texto += f"\033[32m{cell}\033[0m "  # Amigo en verde
            elif cell == '$':
                texto += f"\033[33m{cell}\033[0m "  # Obstáculo en amarillo
            else:
                texto += f"{cell} "
        print(texto)
    print("   \033[34ma b c d e f g h\033[0m")
    return grid

def split_into_chunks(string, chunk_size=3):
    """
    Divide un string en fragmentos del tamaño especificado.
    """
    return [string[i:i + chunk_size] for i in range(0, len(string), chunk_size)]

def find_position(grid, target):
    """
    Encuentra la posición de un carácter objetivo en la cuadrícula.
    Devuelve las coordenadas en formato (letra, número) o None si no se encuentra.
    """
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == target:
                return (number_to_letter[x], abs(y - 8))
    return None

def find_position_obstacle(grid, target):
    """
    Encuentra todas las posiciones de un carácter objetivo en la cuadrícula.
    """
    obstacle_pos = []
    for y, row in enumerate(grid):
        for x, cell in enumerate(row):
            if cell == target:
                column = number_to_letter[x]
                row_number = abs(y - 8)
                obstacle_pos.append((column, row_number))
    return obstacle_pos

def predict_enemy_position(current_pos, target_pos, obstacles):
    """
    Predice la próxima posición de la nave enemiga considerando movimientos ortogonales
    y evitando obstáculos.
    """
    enemy_x, enemy_y = current_pos
    target_x, target_y = target_pos

    # Priorizar movimiento horizontal
    if enemy_x != target_x:
        dx = 1 if enemy_x < target_x else -1  # Determinar dirección en el eje x
        next_pos = (enemy_x + dx, enemy_y)
        if next_pos not in obstacles:
            return next_pos

    # Si no puede moverse horizontalmente, intentar moverse verticalmente
    if enemy_y != target_y:
        dy = 1 if enemy_y < target_y else -1  # Determinar dirección en el eje y
        next_pos = (enemy_x, enemy_y + dy)
        if next_pos not in obstacles:
            return next_pos

    # Si ambos movimientos están bloqueados, quedarse en la posición actual
    return current_pos

number_to_letter = {i: chr(97 + i) for i in range(8)}
letter_to_number = {v: k for k, v in number_to_letter.items()}

def main():
    # Simular llamados al radar
    api_responses = [
        "a01b^1c01d01e01f01g01h01|a02b02c02d$2e02f02g02h02|a03b03c$3d03e03f03g03h03|a04b04c$4d04e04f04g04h04|a05b05c05d05e05f05g05h05|a06b06c06d$6e06f06g06h06|a07b07c07d07e07f07g07h07|a08b08c08d08e#8f08g08h08|",
        "a01b01c01d01e01f01g01h01|a02b02c02d$2e02f02g02h02|a^3b03c$3d03e03f03g03h03|a04b04c$4d04e04f04g04h04|a05b05c05d05e05f05g05h05|a06b06c06d$6e06f06g06h06|a07b07c07d07e07f07g07h07|a08b08c08d08e#8f08g08h08|",
        "a01b01c01d01e01f01g01h01|a02b02c02d$2e02f02g02h02|a03b03c$3d03e03f03g03h03|a04b04c$4d04e04f04g04h04|a05b^5c05d05e05f05g05h05|a06b06c06d$6e06f06g06h06|a07b07c07d07e07f07g07h07|a08b08c08d08e#8f08g08h08|"
    ]

    for i, response in enumerate(api_responses):
        print(f"\nLlamado al radar {i + 1}:")
        grid = parse_grid(response)

        enemy_pos = find_position(grid, '^')
        hope_pos = find_position(grid, '#')
        obstacles = [(letter_to_number[x], y) for x, y in find_position_obstacle(grid, '$')]

        print(f"Posición de la nave enemiga: {enemy_pos}")
        print(f"Posición de la nave amiga: {hope_pos}")

        # Predecir próxima posición del enemigo
        if enemy_pos and hope_pos:
            next_pos = predict_enemy_position((letter_to_number[enemy_pos[0]], enemy_pos[1]), 
                                              (letter_to_number[hope_pos[0]], hope_pos[1]), 
                                              obstacles)
            print(f"Predicción de la próxima posición: {number_to_letter[next_pos[0]]}{next_pos[1]}")

main()
