import re

class WhileCompiler:
    def __init__(self, source_code, symbol_table):
        self.source_code = source_code
        self.symbol_table = symbol_table # Simulasi tabel simbol (variabel yang sudah dideklarasikan)
        self.label_counter = 1
        self.temp_counter = 1

    def new_label(self):
        lbl = f"L{self.label_counter}"
        self.label_counter += 1
        return lbl

    def new_temp(self):
        tmp = f"t{self.temp_counter}"
        self.temp_counter += 1
        return tmp

    def lexical_analysis(self):
        # Spesifikasi token menggunakan Regex
        token_specification = [
            ('KEYWORD',   r'\bwhile\b'),
            ('LPAREN',    r'\('),
            ('RPAREN',    r'\)'),
            ('LBRACE',    r'\{'),
            ('RBRACE',    r'\}'),
            ('OP',        r'[<>=+\-*/]+'),
            ('ID',        r'[A-Za-z_][A-Za-z0-9_]*'),
            ('NUMBER',    r'\d+'),
            ('SKIP',      r'[ \t\n]+'),
            ('MISMATCH',  r'.'),
        ]
        tok_regex = '|'.join('(?P<%s>%s)' % pair for pair in token_specification)
        tokens = []
        
        for mo in re.finditer(tok_regex, self.source_code):
            kind = mo.lastgroup
            value = mo.group()
            if kind == 'SKIP':
                continue
            elif kind == 'MISMATCH':
                raise RuntimeError(f'Karakter tidak valid: {value}')
            tokens.append((kind, value))
            
        return tokens

    def syntax_analysis(self, tokens):
        # Membangun Abstract Syntax Tree (AST) sederhana berbasis dictionary
        ast = {'type': 'WhileLoop'}
        
        try:
            if tokens[0][0] != 'KEYWORD' or tokens[0][1] != 'while':
                raise SyntaxError("Sintaks harus diawali dengan 'while'")
            
            if tokens[1][0] != 'LPAREN':
                raise SyntaxError("Kurang karakter '('")
            
            # Ekstrak Kondisi
            idx = 2
            cond_tokens = []
            while tokens[idx][0] != 'RPAREN':
                cond_tokens.append(tokens[idx])
                idx += 1
            ast['condition'] = cond_tokens
            
            idx += 1
            if tokens[idx][0] != 'LBRACE':
                raise SyntaxError("Kurang karakter '{'")
            
            # Ekstrak Body (Isi Perulangan)
            idx += 1
            body_tokens = []
            while tokens[idx][0] != 'RBRACE':
                body_tokens.append(tokens[idx])
                idx += 1
            ast['body'] = body_tokens
            
            return ast
        except IndexError:
            raise SyntaxError("Struktur sintaks tidak lengkap.")

    def semantic_analysis(self, ast):
        # Mengecek apakah Identifier (ID) yang digunakan sudah ada di Symbol Table
        for kind, val in ast['condition']:
            if kind == 'ID' and val not in self.symbol_table:
                raise NameError(f"Semantic Error: Variabel '{val}' belum dideklarasikan.")
                
        for kind, val in ast['body']:
            if kind == 'ID' and val not in self.symbol_table:
                raise NameError(f"Semantic Error: Variabel '{val}' belum dideklarasikan.")
                
        return True

    def generate_tac(self, ast):
        # Merakit kondisi dan body dari AST menjadi string
        cond_str = " ".join([val for kind, val in ast['condition']])
        body_str = "".join([val for kind, val in ast['body']]) 
        
        label_start = self.new_label()
        label_end = self.new_label()
        
        tac = []
        # Label awal loop
        tac.append(f"{label_start}:")
        
        # Evaluasi kondisi, jika salah (False) lompat ke akhir
        tac.append(f"ifFalse {cond_str} goto {label_end}")
        
        # Eksekusi Body (Simulasi pembuatan variabel temporary untuk operasi hitung)
        if '=' in body_str:
            target, expr = body_str.split('=', 1)
            # Jika berupa operasi matematika (ada +, -, *, /)
            if any(op in expr for op in ['+', '-', '*', '/']):
                temp = self.new_temp()
                tac.append(f"{temp} = {expr}")
                tac.append(f"{target} = {temp}")
            else:
                tac.append(f"{target} = {expr}")
                
        # Lompat kembali ke awal loop
        tac.append(f"goto {label_start}")
        
        # Label akhir loop
        tac.append(f"{label_end}:")
        
        return "\n".join(tac)

    def compile(self):
        print("--- 1. Analisis Leksikal (Tokens) ---")
        tokens = self.lexical_analysis()
        print(tokens)
        
        print("\n--- 2. Analisis Sintaksis (AST) ---")
        ast = self.syntax_analysis(tokens)
        print(ast)
        
        print("\n--- 3. Analisis Semantik ---")
        if self.semantic_analysis(ast):
            print("Status: Sukses. Semua variabel valid terdeklarasi.")
            
        print("\n--- 4. Generasi Three-Address Code (TAC) ---")
        tac = self.generate_tac(ast)
        print(tac)


# ==========================================
# CONTOH PENGGUNAAN
# ==========================================
source_code = "while ( x > 0 ) { x = x - 1 }"
# Anggap 'x' sudah dideklarasikan sebelumnya di baris kode lain
symbol_table = ["x"] 

compiler = WhileCompiler(source_code, symbol_table)
compiler.compile()
