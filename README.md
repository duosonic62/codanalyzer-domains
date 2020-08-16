# Codanalyzer
コードを分析するアプリのドメインモデル

# コード設定
スキーマ設定

```yaml
CodeInputs:
    description: コードの入力値の集合
    type: object
    required: 
    - version
    - codes
    properties:
    version:
      type: string
      description: バージョン情報
      example: '1.0.0'
    codes:
      type: array
      description: コード構造の入力の集合
      items:
        type: object
        required: 
          - name
          - intervals
        properties:
          name:
           type: string
           description: コード名
           example: 'M7'
          intervals:
            type: array
            description: コード構成音
            example: ['R', 'M3', 'P5', 'M7']
            items:
              type: string
              enum: 
                - 'R'
                - 'm2'
                - 'M2'
                - 'm3'
                - 'M3'
                - 'P4'
                - '#4'
                - 'b5'
                - 'P5'
                - '#5'
                - 'm6'
                - 'M6'
                - 'm7'
                - 'M7'
    
```