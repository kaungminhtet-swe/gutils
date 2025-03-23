# Word Counter (`wc`)

## Synopsis

```
wc [OPTION]... [FILE]...
```

## Description

`wc` (Word Counter) is a command-line utility that provides various text statistics for a given file or input. The available options determine which counts are displayed, always in the following order:

- **Newlines**
- **Words**
- **Characters**
- **Bytes**
- **Maximum line length**

## Options

| Option | Long Form | Description               |
| ------ | --------- | ------------------------- |
| `-b`   | `--bytes` | Print the byte count      |
| `-c`   | `--chars` | Print the character count |
| `-l`   | `--lines` | Print the newline count   |
| `-w`   | `--words` | Print the word count      |

## Additional Options

| Option      | Description                                                              |
| ----------- | ------------------------------------------------------------------------ |
| `help`    | Display usage information and exit                                       |
| `version` | Output version details and exit                                          |

## Usage Examples

1. Count lines, words, and characters in a file:
   ```sh
   wc file.txt
   ```
2. Count only lines in a file:
   ```sh
   wc -l file.txt
   ```
3. Count words from standard input:
   ```sh
   echo "Hello, World!" | wc -w
   ```

For more details, refer to the official documentation or use `wc --help`.

## License

Licensed under the Apache License, Version 2.0. You may obtain a copy of the license at:

```
http://www.apache.org/licenses/LICENSE-2.0
```

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
